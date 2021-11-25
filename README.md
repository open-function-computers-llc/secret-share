# OfCO Secret Sharer

Ever wanted to ask a client for a password to their account? Need to get an access code but don't want it stuck in some email log? Use this.

This is a go program. It needs you to build it and then host it. Run the `build.sh` script in the project root, and then put the binary that makes sense for your server out there. Run it as a systemd script, inside of `shell` or however windows works.

The first time you try and build this it might blow up. You will probably need to run `npm install` from the `./frontend` folder before the Vue.js app can be built correctly.
