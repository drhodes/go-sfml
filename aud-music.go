package sfml
// #include <SFML/Audio/Export.h>
// #include <SFML/Audio/SoundStatus.h>
// #include <SFML/Audio/Types.h>
// #include <SFML/Audio/Music.h>
// #include <SFML/System/InputStream.h>
// #include <SFML/System/Time.h>
// #include <SFML/System/Vector3.h>
// #include <stddef.h>
import "C"

import "errors"


type Music struct {
	Cref *C.sfMusic
}

// Create a new music and load it from a file
// This function doesn't start playing the music (call
// Play() to do so).
// Here is a complete list of all the supported audio formats:
// ogg, wav, flac, aiff, au, raw, paf, svx, nist, voc, ircam,
// w64, mat4, mat5 pvf, htk, sds, avr, sd2, caf, wve, mpc2k, rf64.
// \param filename Path of the music file to open
// \return A new sfMusic object (NULL if failed)
func NewMusicFromFile(filename string) (Music, error) { 
	fn := C.CString(filename)
	m := Music{C.sfMusic_createFromFile(fn)}
	if fn == nil {
		return m, errors.New("NewMusicFromFile fails to open: " + filename)
	}
    return m, nil
	// sfMusic* sfMusic_createFromFile(const char* filename);
}


// MEDITATE IF THIS IS FEASIBLE FROM GO. 
// \brief Create a new music and load it from a file in memory
// This function doesn't start playing the music (call
// sfMusic_play to do so).
// Here is a complete list of all the supported audio formats:
// ogg, wav, flac, aiff, au, raw, paf, svx, nist, voc, ircam,
// w64, mat4, mat5 pvf, htk, sds, avr, sd2, caf, wve, mpc2k, rf64.
// \param data        Pointer to the file data in memory
// \param sizeInBytes Size of the data to load, in bytes
// \return A new sfMusic object (NULL if failed)
// sfMusic* sfMusic_createFromMemory(const void* data, size_t sizeInBytes);
// func (self Music) CreateFrommemory(sizeInBytes size_t) *Music { 
//     return C.sfMusic_createFromMemory(self.Cref, sfSize_t(sizeInBytes));
// }


// MEDITATE IF THIS IS FEASIBLE FROM GO. 
// \brief Create a new music and load it from a custom stream
// This function doesn't start playing the music (call
// sfMusic_play to do so).
// Here is a complete list of all the supported audio formats:
// ogg, wav, flac, aiff, au, raw, paf, svx, nist, voc, ircam,
// w64, mat4, mat5 pvf, htk, sds, avr, sd2, caf, wve, mpc2k, rf64.
// \param stream Source stream to read from
// \return A new sfMusic object (NULL if failed)
// sfMusic* sfMusic_createFromStream(sfInputStream* stream);
// func (self Music) Createfromstream() *Music { 
//     return C.sfMusic_createFromStream(self.Cref);
// }           
// \brief Destroy a music
// \param music Music to destroy
// void sfMusic_destroy(sfMusic* music);

func (self Music) Destroy() { 
    C.sfMusic_destroy(self.Cref);
}

// \brief Set whether or not a music should loop after reaching the end
// If set, the music will restart from beginning after
// reaching the end and so on, until it is stopped or
// sfMusic_setLoop(music, sfFalse) is called.
// The default looping state for musics is false.
// \param music Music object
// \param loop  sfTrue to play in loop, sfFalse to play once
// void sfMusic_setLoop(sfMusic* music, sfBool loop);
func (self Music) SetLoop(loop bool) { 
    C.sfMusic_setLoop(self.Cref, Bool(loop));
}
            
// \brief Tell whether or not a music is in loop mode
// \param music Music object
// \return sfTrue if the music is looping, sfFalse otherwise
// sfBool sfMusic_getLoop(const sfMusic* music);
func (self Music) GetLoop() bool { 
    return C.sfMusic_getLoop(self.Cref) == 1
}

// \brief Get the total duration of a music
// \param music Music object
// \return Music duration
// sfTime sfMusic_getDuration(const sfMusic* music);
func (self Music) GetDuration() Time { 
    return Time{C.sfMusic_getDuration(self.Cref)}
}
            
// \brief Start or resume playing a music
// This function starts the music if it was stopped, resumes
// it if it was paused, and restarts it from beginning if it
// was it already playing.
// This function uses its own thread so that it doesn't block
// the rest of the program while the music is played.
// \param music Music object
// void sfMusic_play(sfMusic* music);
func (self Music) Play() { 
    C.sfMusic_play(self.Cref);
}
            
// \brief Pause a music
// This function pauses the music if it was playing,
// otherwise (music already paused or stopped) it has no effect.
// \param music Music object
// void sfMusic_pause(sfMusic* music);
func (self Music) Pause() { 
    C.sfMusic_pause(self.Cref);
}

// \brief Stop playing a music
// This function stops the music if it was playing or paused,
// and does nothing if it was already stopped.
// It also resets the playing position (unlike sfMusic_Pause).
// \param music Music object
// void sfMusic_stop(sfMusic* music);
func (self Music) Stop() { 
    C.sfMusic_stop(self.Cref);
}
            
// \brief Return the number of channels of a music
// 1 channel means a mono sound, 2 means stereo, etc.
// \param music Music object
// \return Number of channels
// unsigned int sfMusic_getChannelCount(const sfMusic* music);
func (self Music) GetChannelCount() uint {
    return uint(C.sfMusic_getChannelCount(self.Cref))
}
            
// \brief Get the sample rate of a music
// The sample rate is the number of audio samples played per
// second. The higher, the better the quality.
// \param music Music object
// \return Sample rate, in number of samples per second
// unsigned int sfMusic_getSampleRate(const sfMusic* music);
func (self Music) GetSampleRate() int { 
    return int(C.sfMusic_getSampleRate(self.Cref))
}

// \brief Get the current status of a music (stopped, paused, playing)
// \param music Music object
// \return Current status
// sfSoundStatus sfMusic_getStatus(const sfMusic* music);
func (self Music) GetStatus() SoundStatus { 
    switch C.sfMusic_getStatus(self.Cref) {
	case 0: return Stopped
	case 1: return Paused
	case 2: return Playing
	}
	return UnknownSoundStatus
}
            
// \brief Get the current playing position of a music
// \param music Music object
// \return Current playing position
// sfTime sfMusic_getPlayingOffset(const sfMusic* music);
func (self Music) GetPlayingOffset() Time { 
    return Time{C.sfMusic_getPlayingOffset(self.Cref)}
}
           
// \brief Set the pitch of a music
// The pitch represents the perceived fundamental frequency
// of a sound; thus you can make a music more acute or grave
// by changing its pitch. A side effect of changing the pitch
// is to modify the playing speed of the music as well.
// The default value for the pitch is 1.
// \param music Music object
// \param pitch New pitch to apply to the music
// void sfMusic_setPitch(sfMusic* music, float pitch);
func (self Music) SetPitch(pitch float32) { 
    C.sfMusic_setPitch(self.Cref, C.float(pitch));
}
            
// \brief Set the volume of a music
// The volume is a value between 0 (mute) and 100 (full volume).
// The default value for the volume is 100.
// \param music  Music object
// \param volume Volume of the music
// void sfMusic_setVolume(sfMusic* music, float volume);
func (self Music) SetVolume(volume float32) { 
    C.sfMusic_setVolume(self.Cref, C.float(volume));
}
           
// \brief Set the 3D position of a music in the audio scene
// Only musics with one channel (mono musics) can be
// spatialized.
// The default position of a music is (0, 0, 0).
// \param music    Music object
// \param position Position of the music in the scene
// void sfMusic_setPosition(sfMusic* music, sfVector3f position);
func (self Music) SetPosition(position Vector3f) { 
    C.sfMusic_setPosition(self.Cref, position.Cref)
}
            
// \brief Make a musics's position relative to the listener or absolute
// Making a music relative to the listener will ensure that it will always
// be played the same way regardless the position of the listener.
// This can be useful for non-spatialized musics, musics that are
// produced by the listener, or musics attached to it.
// The default value is false (position is absolute).
// \param music    Music object
// \param relative sfTrue to set the position relative, sfFalse to set it absolute
// void sfMusic_setRelativeToListener(sfMusic* music, sfBool relative);
func (self Music) SetRelativeToListener(relative bool) { 
    C.sfMusic_setRelativeToListener(self.Cref, Bool(relative));
}
            
// \brief Set the minimum distance of a music
// The "minimum distance" of a music is the maximum
// distance at which it is heard at its maximum volume. Further
// than the minimum distance, it will start to fade out according
// to its attenuation factor. A value of 0 ("inside the head
// of the listener") is an invalid value and is forbidden.
// The default value of the minimum distance is 1.
// \param music    Music object
// \param distance New minimum distance of the music
// void sfMusic_setMinDistance(sfMusic* music, float distance);
func (self Music) SetMinDistance(distance float32) { 
    C.sfMusic_setMinDistance(self.Cref, C.float(distance));
}
            
// \brief Set the attenuation factor of a music
// The attenuation is a multiplicative factor which makes
// the music more or less loud according to its distance
// from the listener. An attenuation of 0 will produce a
// non-attenuated music, i.e. its volume will always be the same
// whether it is heard from near or from far. On the other hand,
// an attenuation value such as 100 will make the music fade out
// very quickly as it gets further from the listener.
// The default value of the attenuation is 1.
// \param music       Music object
// \param attenuation New attenuation factor of the music
// void sfMusic_setAttenuation(sfMusic* music, float attenuation);
func (self Music) SetAttenuation(attenuation float32) { 
    C.sfMusic_setAttenuation(self.Cref, C.float(attenuation));
}
            
// \brief Change the current playing position of a music
// The playing position can be changed when the music is
// either paused or playing.
// \param music      Music object
// \param timeOffset New playing position, in milliseconds
// void sfMusic_setPlayingOffset(sfMusic* music, sfTime timeOffset);
// TODO: file this doc as a bug, the second arg is sfTime, not int.
func (self Music) SetPlayingOffset(time Time) {
	C.sfMusic_setPlayingOffset(self.Cref, time.Cref)
}

// \brief Get the pitch of a music
// \param music Music object
// \return Pitch of the music
// float sfMusic_getPitch(const sfMusic* music);
func (self Music) GetPitch() float32 { 
    return float32(C.sfMusic_getPitch(self.Cref))
}
            
// \brief Get the volume of a music
// \param music Music object
// \return Volume of the music, in the range [0, 100]
// float sfMusic_getVolume(const sfMusic* music);
func (self Music) GetVolume() float32 { 
    return float32(C.sfMusic_getVolume(self.Cref))
}
            
// \brief Get the 3D position of a music in the audio scene
// \param music Music object
// \return Position of the music in the world
// sfVector3f sfMusic_getPosition(const sfMusic* music);
func (self Music) GetPosition() Vector3f { 
    return Vector3f{C.sfMusic_getPosition(self.Cref)}
}

// \brief Tell whether a music's position is relative to the
//        listener or is absolute
// \param music Music object
// \return sfTrue if the position is relative, sfFalse if it's absolute
// sfBool sfMusic_isRelativeToListener(const sfMusic* music);
func (self Music) IsRelativeToListener() bool { 
    return C.sfMusic_isRelativeToListener(self.Cref) == 1
}
            
// \brief Get the minimum distance of a music
// \param music Music object
// \return Minimum distance of the music
// float sfMusic_getMinDistance(const sfMusic* music);
func (self Music) Getmindistance() float32 { 
    return float32(C.sfMusic_getMinDistance(self.Cref))
}
            
// \brief Get the attenuation factor of a music
// \param music Music object
// \return Attenuation factor of the music
// float sfMusic_getAttenuation(const sfMusic* music);
func (self Music) Getattenuation() float32 { 
    return float32(C.sfMusic_getAttenuation(self.Cref))
}
            
