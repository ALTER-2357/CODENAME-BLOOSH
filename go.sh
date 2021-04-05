#!/bin/bash


GOOS=darwin go build . && mkdir mac && mv BLOOSH mac && mkdir mac/cvs && mkdir mac/site

GOOS=windows go build . && mkdir windows && mv BLOOSH.exe windows && mkdir windows/cvs && mkdir windows/site

go build . && mkdir linux && mv BLOOSH linux && mkdir linux/cvs && mkdir linux/site


