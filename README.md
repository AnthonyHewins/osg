OSG (Open source guard)
=======================

# Objective

"Open source" has become a buzzword, and code being open source does not imply
it isn't abusing you, it just means you can find out if it is. You still need to
read it, understand it, and find areas that are troubling. In a repo of millions
of lines of code, this isn't an easy thing to do.

However, there are telling signs to detect such behavior. For example, if
telemetry is taking place, it has to go somewhere. Isolating all network calls
would immediately unveil if such behavior was possible. That's where this tool
comes in. It can get you exactly where the bad parts likely are, and reveal
issues immediately.

## How does it do it?

OSG takes in a repo (a tarball, zip, directory, github repo) and uses regular
expressions to find anything it deems suspicious in a multithreaded manner.
Metrics of how fast it performs are a WIP, but as of right now, ignoring
almost no files whatsoever it's able to go from start to finish in a
few seconds for every codebase I've tested it against.

It will output results in text to STDOUT.

### Results

This list will expand more and more as I use it on more repos but here's the
list so far. 

1. [Mozilla's
   focus-android](https://github.com/mozilla-mobile/focus-android/tree/8d5eea78e7df24ef00b032b838bc4c8ad7688f5d)
   claims to allow the user to "browse like nobody's watching" but has telemetry
   built into it (OSG found [this](https://github.com/mozilla-mobile/focus-android/tree/8d5eea78e7df24ef00b032b838bc4c8ad7688f5d/app/src/main/java/org/mozilla/focus/telemetry)).
   This may be a thing that can be disabled, but at any rate, anyone interesting in auditing the code would likely
   be interested starting here.
