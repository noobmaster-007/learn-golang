# FMT Package
## General
* %v - The value of variable
* %T - The type of the variable
* %% - Print % sign

## Boolean
* %t - bool value (true or false)

## Integer
* %b - binary
* %d - decimal (base10)
* %o - octo
* %x - Hex decimal (base16)

## Floating Points
* %e - scientific notation
* %f, %F - Decimal to exponents
* %g - for large exponents

## Strings
* %s (default)
* %q (double quoted)

## Width & Precision
* %f - default width, default precision
* %9f - width 9, default precision
* %.2f (default width, precision 2)
* %9.2f (width 9, precision 2)
* %9.f (width 9, precision 0)

## Padding
* %09d (pads digit to length 9 with leading 0)
* %-4d (pads with spaces, width 4, left justfied)