# ChillD

*POC: Thermal controller for the Raspberry Pi Compute Module 4 IO Board.*

**ChillD is a simple *user space* thermal controller for the [CM4IO](https://www.raspberrypi.com/products/compute-module-4-io-board/). It periodically...**
- Reads the BCM2711's temperature using the Generic Thermal Sysfs driver of the linux kernel
- Calculates the target fan speed in a [linear](https://github.com/tmsmr/chilld/blob/main/fancurve/linear.go) manner
- Controls a fan connected to the PWM fan connector on the CM4IO board

I mainly built this as an application example for https://github.com/tmsmr/cm4iofan. Please note [should-i-use-this-for-my-247-running-project](https://github.com/tmsmr/cm4iofan#should-i-use-this-for-my-247-running-project).

## Requirements
See https://github.com/tmsmr/cm4iofan#requirements

## Usage

### Get it
- Clone the Repository and build the `chilld` executable using `go build`, or
- Build the `chilld` executable using `go install github.com/tmsmr/chilld@v0.9.0` to find it in `$GOPATH/bin`, or
- Download the latest tagged version from the Release page

### Use it
- No configuration needed (nor available at the moment...)
- Execute it: `./chilld`

## Comments

### System service
- Feel free to adjust and use [chilld.service](./chilld.service)
- Scripts for OpenRC, SysVinit shouldn't be a big deal

### User with access to I2C
ChillD needs a user with access to the I2C bus. If you are using Raspian/RaspiOS, the group `i2c` should be available for that purpose.

Otherwise you may:
- Add a user for ChillD: `useradd -r -M -s /bin/false chilld`
- Add the I2C group: `groupadd i2c`
- Add the user to the created group: `usermod -aG i2c chilld`
- Add a udev rule to assign the group: `echo 'KERNEL=="i2c-[0-9]*", GROUP="i2c"' >> /etc/udev/rules.d/10-i2c_group.rules`
