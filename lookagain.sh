#! bin/bash

find . -name '*.sh' | sed 's/.sh/\t/g' cut -f 1 | -d '.' -f 2 | cut -d '/' -f 2


