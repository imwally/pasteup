Pasteup
=======

Create anonymous gists directly from the command line.

<pre><code>$pasteup README.md 
https://gist.github.com/601f1e1ca03cb79dad03
</code></pre>

If no arguments are given then vi is launched to create a new paste. After pasting, just write and close vi (:wq) and pasteup will return the gist url.

TODO
----
* Launch default editor of environment instead of launching vi.
* Add authentication for user gists.
* Add raw view URLs to response. 
