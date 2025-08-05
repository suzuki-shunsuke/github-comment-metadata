# github-comment-metadata

[![Go Reference](https://pkg.go.dev/badge/github.com/suzuki-shunsuke/github-comment-metadata.svg)](https://pkg.go.dev/github.com/suzuki-shunsuke/github-comment-metadata)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/github-comment-metadata.svg)](https://github.com/suzuki-shunsuke/github-comment-metadata)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/github-comment-metadata/master/LICENSE)

Parser of metadata [github-comment](https://github.com/suzuki-shunsuke/github-comment) embeds into comment.

## Background

github-comment is a CLI tool to post a comment to GitHub commit, issue and pull request.
From v3, github-comment embeds metadata into comment to support hide comment.

* [github-comment hide](https://github.com/suzuki-shunsuke/github-comment#hide)
* [github-comment#210](https://github.com/suzuki-shunsuke/github-comment/pull/210)

In addition to github-comment, we develop [tfcmt](https://github.com/suzuki-shunsuke/tfcmt), which is a fork of [mercari/tfnotify](https://github.com/mercari/tfnotify) and post a comment like github-comment.
To embed metadata into comment via tfcmt like github-comment, we decide to release this feature as a library and use this in github-comment and tfcmt.

## metadata format

```
<!-- github-comment: JSON object -->
```

ex.

```
<!-- github-comment: {"JobID":"xxx","JobName":"plan","SHA1":"79acc0778da6660712a65fd41a48b72cb7ad16c0","TemplateKey":"default","Vars":{}} -->
```

## License

[MIT](LICENSE)
