core
=====

Mono-repo that houses prtech.news core functionality


### Private Repo Usage

Connecting to a private repo in Go is not complex. However, depending on your machine's setup there could be authentication issues.

1. To avoid such issues make sure you create an SSH key and have it associated to your github account. 
2. Verify your machine authenticates via SSH and the above key by running `ssh -Tv git@github.com`
```shell
Hi your-github-username! You've successfully authenticated, but GitHub does not provide shell access.
debug1: channel 0: free: client-session, nchannels 1
Transferred: sent 2044, received 2552 bytes, in 0.1 seconds
Bytes per second: sent 19605.0, received 24477.5
debug1: Exit status 1
```
3. Setup your `~/.gitconfig`'s content to look like the below:
```shell
[user]
        name = Jose Diaz
        email = ***REMOVED***
```
After these steps you should be good to. Below are optional steps:
4. (Optional) Create a github personal access
5. (Optional) Set up your `~$HOME/.netrc`: (Interpolate the GITHUB_TOKEN)
```shell
machine git@github.com

login your-github-username
password ${GITHUB_TOKEN}
```