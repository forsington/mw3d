# mw3d
Milky Way 3D Visualiztion of GDR3 data.

## Using
* [g3n](https://github.com/g3n/engine) as engine
* [Spacescape](http://alexcpeterson.com/spacescape/) for Skybox

## Todos
- [x] g3n poc
- [ ] Use Effective temperature (Teff) set (161,497,595 stars) and convert to [RGB color data](https://astronomy.stackexchange.com/questions/46664/straightforward-conversion-from-gaia-gdr3-photometry-to-rgb-values), and [this thread](https://stackoverflow.com/questions/21977786/star-b-v-color-index-to-apparent-rgb-color)
    - [ ] [lib for temperature conversion](http://www.vendian.org/mncharity/dir3/blackbody/UnstableURLs/bbr_color.html)
- [ ] Download script for [raw csv files](http://cdn.gea.esac.esa.int/Gaia/gdr3/) (~10 TB)
- [ ] Script to prune data to only stars with parallax measurements (and maybe other criteria)
- [ ] convert galactic coordinates to cartesian coordinates, docs [here] (https://chandra.si.edu/build/navigation.html) and [here](https://astronomy.stackexchange.com/questions/35536/transforming-galactic-coordinates-in-cartesian-with-distance)

## Docs
* [Gaia 3 docs](https://www.cosmos.esa.int/web/gaia/dr3) and [data](cdn.gea.esac.esa.int/Gaia/gdr3)
