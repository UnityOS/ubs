# UBS Package Format

UBS stands for "UnityOS Build System". It is a `golang` application that has support
for building UnityOS Packages.

## UBS Structure
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
        
        }
        
        
        library "requiredLib" {
            minVersion = "1.0.0"
            maxVersion = "2.0.0"
        }
        
    }
    
    // Actual build steps
    build {
    
        // use "externalPlugin" {
        // }
        // RESERVED FOR FUTURE USE
   
        arch "amd64" {
            type "generic" {
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
Contents folder contains root of the source code