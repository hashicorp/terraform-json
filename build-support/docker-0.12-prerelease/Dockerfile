FROM golang

ARG GPG_KEY="51852D87348FFC4C"
ARG BUILD_REF="master"

RUN sh -c "\
  apt-get update && \
  apt-get -y install git curl gpg unzip && \
  mkdir -p /providers-download /build /go/bin /root/.gnupg && \
  chmod og-wrxX -R /root/.gnupg && \
  echo disable-ipv6 > /root/.gnupg/dirmngr.conf && \
  gpg --no-tty --recv-keys ${GPG_KEY} && \
  echo '5\ny\n' | gpg --no-tty --command-fd 0 --edit-key ${GPG_KEY} trust && \
  cd /providers-download && \
  curl -O https://releases.hashicorp.com/terraform/0.12.0-alpha4/terraform_0.12.0-alpha4_terraform_0.12.0-alpha4_linux_amd64.zip && \
  curl -O https://releases.hashicorp.com/terraform/0.12.0-alpha4/terraform_0.12.0-alpha4_SHA256SUMS && \
  curl -O https://releases.hashicorp.com/terraform/0.12.0-alpha4/terraform_0.12.0-alpha4_SHA256SUMS.sig && \
  gpg --verify terraform_0.12.0-alpha4_SHA256SUMS.sig && \
  shasum -c terraform_0.12.0-alpha4_SHA256SUMS 2>/dev/null | grep 'terraform_0.12.0-alpha4_terraform_0.12.0-alpha4_linux_amd64.zip: OK' && \
  unzip terraform_0.12.0-alpha4_terraform_0.12.0-alpha4_linux_amd64.zip -x terraform -d /go/bin && \
  cd /build && \
  git clone https://github.com/hashicorp/terraform.git terraform && \
  cd terraform && \
  git checkout ${BUILD_REF} && \
  go install -mod=vendor ./ && \
  rm -rf /providers-download /build && \
  apt-get -y remove git curl gpg unzip && \
  apt-get -y autoremove && \
  rm -rf /var/cache/apt"

COPY entrypoint.sh /

ENTRYPOINT ["/entrypoint.sh"]
