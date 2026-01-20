# Setting up Raspberry Pi 5 for e-Paper (From Scratch)

Since your Pi 5 is brand new, you need to install an Operating System on an SD card first.

## 1. Flash the OS (On your Mac)

1.  Download and install **[Raspberry Pi Imager](https://www.raspberrypi.com/software/)**.
2.  Insert your microSD card into your Mac.
3.  Open Raspberry Pi Imager:
    - **Device**: Select **Raspberry Pi 5**.
    - **OS**: Select **Raspberry Pi OS (64-bit)**.
    - **Storage**: Select your SD card.
4.  Click **Next**. It will ask to apply "OS Customisation" settings. **CLICK EDIT SETTINGS**:
    - **Hostname**: `raspberrypi`
    - **Username/Password**: Create one (e.g., `pi` / `raspberry`).
    - **Wireless LAN**: Enter your WiFi name and password so it connects automatically.
    - **Services**: Enable **SSH** (Use password authentication).
5.  Click **YES** / **SAVE** and write the card.

## 2. Enable SPI (Before ejecting!)

Once the writing finishes, the SD card might unmount. **Re-insert it into your Mac.** You should see a drive named `bootfs`.

1.  Open the `bootfs` drive in Finder.
2.  Find the file named `config.txt`.
3.  Open it with a text editor.
4.  Scroll to the bottom and ensure this line exists (add it if missing):
    ```
    dtparam=spi=on
    ```
5.  Save and Eject.

## 3. Boot & Run

1.  Insert SD card into Pi 5.
2.  Power it on. Wait 2-3 minutes for first boot.
3.  Open Terminal on your Mac and connect:
    ```bash
    ssh pi@raspberrypi.local
    ```
    (Type "yes" if asked, then your password).

## 4. Test Display

Once logged in via SSH, run the test commands:

```bash
sudo apt update && sudo apt install -y git python3-pip python3-pil python3-numpy python3-lgpio

git clone https://github.com/waveshare/e-Paper
cd e-Paper/RaspberryPi_JetsonNano/python
sudo python3 setup.py install
cd examples
python3 epd_2in13_V2_test.py
```
