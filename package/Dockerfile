FROM registry.suse.com/bci/bci-base:15.4
RUN zypper -n rm container-suseconnect && \
    zypper install -y iptables=1.8.7 && \
    zypper -n clean -a && rm -rf /tmp/* /var/tmp/* /usr/share/doc/packages/*
COPY bin/cloudweav-network-controller /usr/bin/
CMD ["cloudweav-network-controller"]
