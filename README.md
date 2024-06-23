# psfv: Estimate Baseball/Softball Pitch Speed from Video Clip

`psfv` (Pitching Speed From Video) is a simple tool to estimate the speed of a
pitched softball or baseball from a video clip (i.e., one that was taken using a
cell phone's camera), without using more accurate equipment such as
[PITCHf/x](https://en.wikipedia.org/wiki/PITCHf/x) or a radar gun.

## The methodology

The `psfv` tool computes the __average velocity__ `v` of a pitched ball, using
the simple formula 

```text
v = d / t
```

where `d` and `t` are the distance and duration, respectively, from the ball
leaves a pitcher's hand until it is above the home plate. Note that our estimate
is different from and slower than the MLB [velocity
(VELO)](https://www.mlb.com/glossary/statcast/velocity), which is defined to be
"the maximum speed of a given pitch at any point from its release to the time it
crosses home plate" (effectively the speed of the ball at the release point due
to physics).

## How to use

Currently, only softball pitch speed estimation is supported. Baseball support
will be added in the future.

### Install dependency and `psfv`

`psfv` depends on `ffprobe`, which can be obtained by installing `ffmpeg`.

On macOS

```bash
brew install ffmpeg
```

On Ubuntu 24.04 LTS

```bash
sudo apt install ffmpeg
```

Install `psfv`

```bash
go install github.com/syncom/psfv
```

### Prepare the video clip

Before AI can help us do better, one needs to do the video clip preparation step manually. The manual work here isn't too difficult, yet needs to be performed carefully.

To obtain the input video clip for `psfv`, load your game footage into a video
editor supporting video trimming at the frame level. For example, [QuickTime
Player for
Mac](https://support.apple.com/guide/quicktime-player/trim-a-movie-or-clip-qtpf2115f6fd/mac)
can be used to do this. Carefully trim the video so that the resulting clip
starts at the frame when the ball is released from the pitcher's hand, and end
at the frame when the ball crosses the home plate. The duration of the video
clip should be less than 1 second in practical scenarios. A sample prepared
video clip is [pitch_sample_10u.mp4](./testdata/pitch_sample_10u.mp4)

### Run `psfv`

For softball 10U

```bash
$ psfv softball --age-group 10u testdata/pitch_sample_10u.mp4 
Duration of video: 0.636003
Age group: 10u
Average pitch speed: 33.233119 mph (53.483388 kph)
```

Please refer to `psfv --help` and `psfv softball --help` for more usage detail.
