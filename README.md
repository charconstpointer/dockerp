# dockerp
### How it works
For each docker container running dockerp will create a tcp tunnel routing your localhost traffic directly to your remote machine hosting docker engine instance.
### Example
Let's say you have linux vm running with docker installed on it and an M1 Apple laptop

You can't run some of your favorites images on M1

But you can run them on your linux VM! 🥳

Expose docker engine on your linux machine
>sudo dockerd -H unix:///var/run/docker.sock -H tcp://0.0.0.0:2375

Setup docker host for your M1 docker instance
>export DOCKER_HOST=tcp://ip_of_your_linux_machine:2375

Run dockerp with flag -ds ip_of_your_linux_machine

>go run main.go -ds 18.170.38.77 <-- this address will change for you

Now any docker image you run with docker will be hosted on your remote docker engine instance, but will be available to access on localhost

>docker run -p 80:80 -d nginxdemos/hello
Is now accessible on your localhost M1
