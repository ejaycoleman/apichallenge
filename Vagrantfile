# -*- mode: ruby -*-
# vi: set ft=ruby :
Vagrant.configure("2") do |config|
  config.vm.define "manager1" do |manager|
    manager.vm.box = "williamyeh/ubuntu-trusty64-docker"
    manager.vm.hostname = "manager1"
    manager.vm.network "private_network", :type => 'dhcp', :dhcp_ip => '192.168.56.100', :dhcp_lower => '192.168.56.101', :dhcp_upper => '192.168.56.254', :adapter => 2
  end

  config.vm.define "worker1" do |worker|
    worker.vm.box = "williamyeh/ubuntu-trusty64-docker"
    worker.vm.hostname = "worker1"
    config.vm.network "private_network", :type => 'dhcp', :dhcp_ip => '192.168.56.100', :dhcp_lower => '192.168.56.101', :dhcp_upper => '192.168.56.254', :adapter => 2
  end

  config.vm.define "worker2" do |worker|
    worker.vm.box = "williamyeh/ubuntu-trusty64-docker"
    worker.vm.hostname = "worker2"
    config.vm.network "private_network", :type => 'dhcp', :dhcp_ip => '192.168.56.100', :dhcp_lower => '192.168.56.101', :dhcp_upper => '192.168.56.254', :adapter => 2
  end
end