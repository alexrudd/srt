# srt
Versatile sort tool - use regex to select chunks of text and sort them by specific subsections

I hacked this together to sort the variables.tf files in my Terraform modules.

If that's what you want, then this worked for me:

```bash
srt -file=variables.tf \
  -match='(^|\n)(variable\s+?"(.+?)"\s+?{[\S\s]*?})' \
  -sort-on='$3' \
  -out='$2\n\n'
```

That matches all variable declarations, sorts them by the 3rd matched segment `"(.+?)"`, and outputs the second matched segment `(variable\s+?"(.+?)"\s+?{[\S\s]*?})` followed by two newlines.


## TODO

* Be able to pipe content to it
* better handling of errors
* tests would be nice
* use one of the existing go cli frameworks
