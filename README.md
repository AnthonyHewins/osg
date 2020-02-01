## What does this do?

If you use open source, how can you know it's acting the way you'd like it to?
You'd have to read the whole thing, or at least to the part where something
fishy or malicious starts happening. Doing so is nearly impossible today, who
knows how many millions of lines of code you'd need to read to find exactly
where something goes wrong.

However, there are telling signs that you can engineer a way to detect such
behavior. For example, if telemetry is taking place, it has to go somewhere.
Finding an instance of that would immediately give you a starting point to audit
the code; that's where this tool comes in. It can get you exactly where the bad
parts likely are, and you can have an opinion about the codebase **incredibly**
faster than reading it by yourself.

## How does it do it?

OSG takes in a repo (a tarball, zip, directory, github repo) and uses regular
expressions to find anything it deems suspicious in a multithreaded manner.
Metrics of how fast it performs are a WIP, but as of right now, ignoring
almost no files whatsoever it's able to go from downloading to completion in a
few seconds for every codebase I've tested it against.

### Proof of it working

This list will expand more and more as I use it on more repos but here's the
list so far. 

1. [Mozilla's
   focus-android](https://github.com/mozilla-mobile/focus-android/tree/8d5eea78e7df24ef00b032b838bc4c8ad7688f5d)
   claims to allow the user to "browse like nobody's watching" but has telemetry
   built into it (OSG found [this](https://github.com/mozilla-mobile/focus-android/tree/8d5eea78e7df24ef00b032b838bc4c8ad7688f5d/app/src/main/java/org/mozilla/focus/telemetry)). Whether or not this can be enabled or
   disabled is something I haven't looked into yet, so don't damn Mozilla yet.
   
