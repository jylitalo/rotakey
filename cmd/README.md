# github.com/jylitalo/rotakey/cmd

## Overview
Package cmd provides spf13/cobra interface to package

Imports: 6

## Index
- [func NewCmd(optFns ...func(*Options)) *cobra.Command](#func-newcmd)
- [type Options](#type-options)

## Examples

This section is empty.

## Constants

This section is empty.

## Variables
This section is empty.

## Functions

### func [NewCmd](./cmd.go#L21)

<pre>
func NewCmd(optFns ...func(<a href="#type-options">*Options</a>)) <a href="https://pkg.go.dev/github.com/spf13/cobra#Command">*cobra.Command</a>
</pre>
## Types
### type [Options](./cmd.go#L14)

<pre>
type Options struct {
    Use string
    AwsCfg <a href="../types/README.md#type-awsconfig">types.AwsConfig</a>
    FileCfg <a href="../types/README.md#type-dotaws">types.DotAws</a>
    Rotate <a href="../types/README.md#type-rotate">types.Rotate</a>
}
</pre>

--

Generated by [github.com/jylitalo/go2md](https://github.com/jylitalo/go2md/) v0.4.1

