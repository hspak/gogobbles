# gogobbles
### Turkey themed todo app

## URL

[Gogobbles](https://gogobbles.com)!

## What is it?

Consists of 2 pages:

    /

Landing page,

    /list/:label

Your list.

## Is it realtime?

It's kind of. It'll update every 3 seconds and add new items that were added for
anyone on the same page. Removed items don't get updated unless they refresh the
page.

## Stack

Server - Go - Martini

DB - MongoDB - mgo

Frontend - Bare JS

Future:

I plan on removing Martini as a dependency and picking up ReactJS to replace my
hacked javascript.

## API

Docs [here](https://gogobbles.com/api)

## LICENSE

MIT
