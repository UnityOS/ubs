# UBS Package Format

Version 0.0.1

UBS stands for "UnityOS Build System". It is a `golang` application that has support
for building UnityOS Packages.

## UBS Source Package Structure
```text
package root
 |
 ------- package.hcl
 |
 ------- Contents
```

### package.hcl
File providing instructions for the build system how to build a package and what is needed.
It is largely inspired by Hashicorp Waypoint.

```hcl
// 
name = "Package name"
version = "0.1.0" // According to Semantic Versioning
maintainer = ""

checkout {} // RESERVED FOR FUTURE USE

build {
    placeholder = false // mark true if i.e. ISO should be created
    // Defines prerequisites - for now informational only.
    prerequisites {
    
        // libraries and packages are carried over to the built package
        // so in this example built product would require 
        // requiredLib and meson as a dependencies
        
        tool "make" {
            minVersion = "1.0.0"
            maxVersion = "2.0.0"
        }    
        
        package "meson" {
            minVersion = "1.0.0"
            maxVersion = "2.0.0"
        }
        
        libraryDevel "requiredLib-dev" {
            // This library is required only on the development / build workstation
            // And not required on the target
        }
        
        
        library "requiredLib" {
            minVersion = "1.0.0"
            maxVersion = "2.0.0"
        }
        
    }
    
    // Actual build steps
    build {
        // Package is required to be built on same type host
        // i.e. False could be for golang
        sameArch = false
        
        // use "externalPlugin" {
        // }
        // RESERVED FOR FUTURE USE
   
        arch "amd64" {
            type "generic" {
                env {
                    // environment variables (if needed)
                }
                
                commands = [
                    "./configure",
                    "make",
                    "make install"
                ]
            }
        }
        
        arch "aarch64" {
            type "generic" {
                commands = [
                    "./configure",
                    "make",
                    "make install"
                ]
            }
        }
    }
    
    // Defines outcome for `upt` packager.
    products {
        file "library.so" {
            source = "./library.so"
            destination = "/usr/local/lib/library.so"
        }
        file "library2.so" {
            source = "./library2.so"
            destination = "/usr/local/lib/library2.so"
        }
    }
}

```
### Contents
Contents folder contains root of the source code.

## Format
The `package.hcl` file and the Contents folder is `tar`-ed into an archive, and then
compressed with zstd. 

## Digital Signature
TBA - not yet supported

## UBS Binary Package Structure

```text
package root
 |
 ------- package.hcl
 |
 ------- Contents
```

### package.hcl
File providing instructions for the build system how to build a package and what is needed.
It is largely inspired by Hashicorp Waypoint.

```hcl
// 
name = "Package name"
version = "0.1.0" // According to Semantic Versioning
maintainer = ""

contents{}
requirements{}
preinstall{}
install{}
postinstall {}
preupdate{}
postupdate{}
preremove{}
remove{}
postremove{}
```
### Contents
Contents folder contains distribution of the package.

## Format
The `package.hcl` file and the Contents folder is `tar`-ed into an archive, and then
compressed with zstd.

## Digital Signature
TBA - not yet supported