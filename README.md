# MemeLord79

MemeLord79 is a Meme Lord bent on posting images to Slack. 
Originally, the application was going to only be built into Github's Actions,
but then I changed my mind because a lot of prepwork needs to be done in the 
container to save on time, so I might as well just use Go and package it into
a stratch container.


## Size Matters

One of the thing that I've noticed, is that at the time of writing the 
application is slightly bigger than I'd like it to be, and I'm wondering if 
using another language could shrink down my footprint so the container gets
loaded into actions faster to save on build minutes. Another thing to contend
with is that we need to run this as fast as possible, so that way we don't take
up build minutes.


## go-git vs apk git

go-get is smaller than regular git. It increases the code footprint, but 
decreases the container size in the end.


## Authentication

Generate an RSA keypair. 

Put the public key in the git repo's deploy token section, and enable 
"Allow write access". This will allow MemeLord to use the git repository and 
use the "state" branch as a pseudo database.

The private key should be in the env `SSH_PRIVATE_KEY`. 


## Docker

The docker image provided will build using alpine, and then use scratch as the 
running image. UPX is used to make the binary smaller, and to help with making
the docker container more compact and faster to move across the wire.