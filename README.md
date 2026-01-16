# github-comment-metadata

[![DeepWiki](https://img.shields.io/badge/Ask_DeepWiki-000000.svg?logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACwAAAAyCAYAAAAnWDnqAAAAAXNSR0IArs4c6QAAA05JREFUaEPtmUtyEzEQhtWTQyQLHNak2AB7ZnyXZMEjXMGeK/AIi+QuHrMnbChYY7MIh8g01fJoopFb0uhhEqqcbWTp06/uv1saEDv4O3n3dV60RfP947Mm9/SQc0ICFQgzfc4CYZoTPAswgSJCCUJUnAAoRHOAUOcATwbmVLWdGoH//PB8mnKqScAhsD0kYP3j/Yt5LPQe2KvcXmGvRHcDnpxfL2zOYJ1mFwrryWTz0advv1Ut4CJgf5uhDuDj5eUcAUoahrdY/56ebRWeraTjMt/00Sh3UDtjgHtQNHwcRGOC98BJEAEymycmYcWwOprTgcB6VZ5JK5TAJ+fXGLBm3FDAmn6oPPjR4rKCAoJCal2eAiQp2x0vxTPB3ALO2CRkwmDy5WohzBDwSEFKRwPbknEggCPB/imwrycgxX2NzoMCHhPkDwqYMr9tRcP5qNrMZHkVnOjRMWwLCcr8ohBVb1OMjxLwGCvjTikrsBOiA6fNyCrm8V1rP93iVPpwaE+gO0SsWmPiXB+jikdf6SizrT5qKasx5j8ABbHpFTx+vFXp9EnYQmLx02h1QTTrl6eDqxLnGjporxl3NL3agEvXdT0WmEost648sQOYAeJS9Q7bfUVoMGnjo4AZdUMQku50McDcMWcBPvr0SzbTAFDfvJqwLzgxwATnCgnp4wDl6Aa+Ax283gghmj+vj7feE2KBBRMW3FzOpLOADl0Isb5587h/U4gGvkt5v60Z1VLG8BhYjbzRwyQZemwAd6cCR5/XFWLYZRIMpX39AR0tjaGGiGzLVyhse5C9RKC6ai42ppWPKiBagOvaYk8lO7DajerabOZP46Lby5wKjw1HCRx7p9sVMOWGzb/vA1hwiWc6jm3MvQDTogQkiqIhJV0nBQBTU+3okKCFDy9WwferkHjtxib7t3xIUQtHxnIwtx4mpg26/HfwVNVDb4oI9RHmx5WGelRVlrtiw43zboCLaxv46AZeB3IlTkwouebTr1y2NjSpHz68WNFjHvupy3q8TFn3Hos2IAk4Ju5dCo8B3wP7VPr/FGaKiG+T+v+TQqIrOqMTL1VdWV1DdmcbO8KXBz6esmYWYKPwDL5b5FA1a0hwapHiom0r/cKaoqr+27/XcrS5UwSMbQAAAABJRU5ErkJggg==)](https://deepwiki.com/suzuki-shunsuke/github-comment-metadata)
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
