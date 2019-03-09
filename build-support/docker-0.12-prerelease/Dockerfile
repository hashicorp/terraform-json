FROM golang

# AWS Provider
ARG AWS_BUILD_REF="v0.12-sdk-updates-b758628e51df"
RUN sh -c "\
  apt-get update && \
  apt-get -y install git && \
  mkdir -p /build && \
  cd /build && \
  git clone https://github.com/terraform-providers/terraform-provider-aws.git terraform-provider-aws && \
  cd terraform-provider-aws && \
  git checkout ${AWS_BUILD_REF} && \
  go install -mod=vendor ./ && \
  rm -rf /build && \
  apt-get -y remove git && \
  apt-get -y autoremove && \
  rm -rf /var/cache/apt"

# Null Provider
ARG NULL_BUILD_REF="v0.12-sdk-updates-b758628e51df"
RUN sh -c "\
  apt-get update && \
  apt-get -y install git && \
  mkdir -p /build && \
  cd /build && \
  git clone https://github.com/vancluever/terraform-provider-null.git terraform-provider-null && \
  cd terraform-provider-null && \
  git checkout ${NULL_BUILD_REF} && \
  go install -mod=vendor ./ && \
  rm -rf /build && \
  apt-get -y remove git && \
  apt-get -y autoremove && \
  rm -rf /var/cache/apt"

# Random Provider
ARG RANDOM_BUILD_REF="v0.12-sdk-updates-b758628e51df"
RUN sh -c "\
  apt-get update && \
  apt-get -y install git && \
  mkdir -p /build && \
  cd /build && \
  git clone
  https://github.com/vancluever/terraform-provider-random.git terraform-provider-random && \
  cd terraform-provider-random && \
  git checkout ${RANDOM_BUILD_REF} && \
  go install -mod=vendor ./ && \
  rm -rf /build && \
  apt-get -y remove git && \
  apt-get -y autoremove && \
  rm -rf /var/cache/apt"

# Core
ARG CORE_BUILD_REF="e466bc4bb73318de7f2b63d9f99e584d35117379"
RUN sh -c "\
  apt-get update && \
  apt-get -y install git && \
  mkdir -p /build && \
  cd /build && \
  git clone https://github.com/hashicorp/terraform.git terraform && \
  cd terraform && \
  git checkout ${BUILD_REF} && \
  go install -mod=vendor ./ && \
  rm -rf /build && \
  apt-get -y remove git && \
  apt-get -y autoremove && \
  rm -rf /var/cache/apt"

COPY entrypoint.sh /

ENTRYPOINT ["/entrypoint.sh"]
