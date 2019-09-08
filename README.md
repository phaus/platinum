# platinum

Insert Plantuml Graphs into Markdow.

## Build

> docker build -t phaus/platinum .

## RUN

Create SVG Images:

> docker run -v $PWD:/project phaus/platinum:latest --kind svg --input /project/testdata/test1.md --output /project/build/

Create ASCII ART:

>  docker run -v $PWD/testdata:/testdata -v $PWD:/build:/build  phaus/platinum:latest --kind txt --input testdata/test1.md --output /build/
