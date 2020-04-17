# MemerLord79

MemerLord79 is a Meme Lord bent on posting images to Slack. 
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