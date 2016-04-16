# gdc
GitHub release download counter

### Instalation
gdc is single binary program.
Download and deploy any directory.

And you must setting environment variables below.

| Name | desc |
| GDC_ID | GitHub account id |
| GDC_ACCESS_TOKEN | GitHub API Access Token |

GitHub Access token get from
```
Home -> Setting -> Personal access tokens -> Generate new token
```

### Usage

```
gdc [your repository name]
```

### Example

```
$gdc hbp

[Name]: hbp
[Total]: 2
**************************************
[Tag Name]: 0.1
[Release Name]: 0.1
[Total Downloads]: 2
--------------------------------------
[Name]: darwin_386.zip
[Download Count]: 0
[Name]: darwin_amd64.zip
[Download Count]: 0
[Name]: linux_386.zip
[Download Count]: 0
[Name]: linux_amd64.zip
[Download Count]: 1
[Name]: windows_386.zip
[Download Count]: 0
[Name]: windows_amd64.zip
[Download Count]: 1
```


### Reference

[GitHub API - Releases](https://developer.github.com/v3/repos/releases/)
