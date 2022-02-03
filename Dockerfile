ARG BUILD_USER="build"
ARG SERVICE_USER="service"
ARG BUILD_DEPENDENCIES="make go"
ARG SERVICE_DEPENDENCIES="openrc"
ARG BUILD_DIRECTORY="/home/${BUILD_USER}/service"
ARG BUILD_OUT_DIRECTORY="${BUILD_DIRECTORY}/out"
ARG SERVICE_DIRECTORY="/home/${SERVICE_USER}"
ARG EXECUTABLE="google_dynamic_dns_update_service"
ARG EXECUTABLE_PATH="/usr/local/bin"

FROM alpine as base

RUN apk update && apk upgrade --available --no-cache --prune --purge


FROM base as build_prep
ARG BUILD_USER
ARG BUILD_DIRECTORY
ARG BUILD_DEPENDENCIES

RUN apk add ${BUILD_DEPENDENCIES} && adduser -D "${BUILD_USER}"

USER "${BUILD_USER}"

WORKDIR "${BUILD_DIRECTORY}"

COPY --chown="${BUILD_USER}" . .

FROM build_prep as build

RUN make


FROM alpine as openrc_prep
ARG SERVICE_USER
ARG EXECUTABLE
ARG EXECUTABLE_PATH
ARG SERVICE_DIRECTORY
ARG BUILD_OUT_DIRECTORY
ARG SERVICE_DEPENDENCIES

ENV EXECUTABLE_NAME=${EXECUTABLE}

RUN apk add ${SERVICE_DEPENDENCIES} && adduser -D "${SERVICE_USER}" \
  && mkdir /run/openrc && touch /run/openrc/softlevel && rc-status

WORKDIR "${SERVICE_DIRECTORY}"

COPY --from=build --chown=root "${BUILD_OUT_DIRECTORY}/${EXECUTABLE}" "${EXECUTABLE_PATH}"
COPY --chown=root "openrc/service" "/etc/init.d/${EXECUTABLE}"

ENTRYPOINT [ "sh", "-c", "rc-service ${EXECUTABLE_NAME} start" ]


FROM openrc_prep as openrc
