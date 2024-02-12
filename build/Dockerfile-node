FROM fedora:39 AS build-env
WORKDIR /go/src/github.com/rusik69/govnocloud
COPY ../ ./
RUN dnf -y install make go libvirt-devel make bash gcc git
RUN make get
RUN make build
FROM fedora:39
WORKDIR /
RUN dnf -y install --quiet --setopt=install_weak_deps=False libvirt ca-certificates qemu-img guestfs-tools
RUN dnf clean all
COPY --from=build-env /go/src/github.com/rusik69/govnocloud/bin/govnocloud-linux-amd64 /usr/local/bin/
RUN chmod 755 /usr/local/bin/govnocloud-linux-amd64
EXPOSE 6969