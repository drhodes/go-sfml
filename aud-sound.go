package sfml

// #cgo LDFLAGS:-lcsfml-audio
// #include <SFML/Audio/SoundStatus.h>
// #include <SFML/System/Time.h>
// #include <SFML/System/Vector3.h>
// #include <SFML/Audio/Sound.h>
import "C"

type SoundStatus uint

const (
	Stopped SoundStatus = iota // Sound / music is not playing
	Paused                     // Sound / music is paused
	Playing                    // Sound / music is playing
	UnknownSoundStatus
)

type Sound struct {
	Cref *C.sfSound
}

// \brief Create a new sound
// \return A new sfSound object
// sfSound* sfSound_create(void);
func (self Sound) NewSound() Sound {
	return Sound{C.sfSound_create()}
}

// \brief Create a new sound by copying an existing one
// \param sound Sound to copy
// \return A new sfSound object which is a copy of \a sound
// sfSound* sfSound_copy(sfSound* sound);
func (self Sound) Copy() Sound {
	return Sound{C.sfSound_copy(self.Cref)}
}

// \brief Destroy a sound
// \param sound Sound to destroy
// void sfSound_destroy(sfSound* sound);
func (self Sound) Destroy() {
	C.sfSound_destroy(self.Cref)
}

// \brief Start or resume playing a sound
// This function starts the sound if it was stopped, resumes
// it if it was paused, and restarts it from beginning if it
// was it already playing.// This function uses its own thread so that it doesn't block
// the rest of the program while the sound is played.
// \param sound Sound object
// void sfSound_play(sfSound* sound);
func (self Sound) Play() { 
    C.sfSound_play(self.Cref);
}

// \brief Pause a sound
// This function pauses the sound if it was playing,
// otherwise (sound already paused or stopped) it has no effect.
// \param sound Sound object
// void sfSound_pause(sfSound* sound);
func (self Sound) Pause() {
	C.sfSound_pause(self.Cref)
}

// \brief Stop playing a sound// This function stops the sound if it was playing or paused,
// and does nothing if it was already stopped.
// It also resets the playing position (unlike sfSound_pause).
// \param sound Sound object
// void sfSound_stop(sfSound* sound);
func (self Sound) Stop() {
	C.sfSound_stop(self.Cref)
}

// \brief Set the source buffer containing the audio data to play
// It is important to note that the sound buffer is not copied,
// thus the sfSoundBuffer object must remain alive as long
// as it is attached to the sound.
// \param sound  Sound object
// \param buffer Sound buffer to attach to the sound
// void sfSound_setBuffer(sfSound* sound, const sfSoundBuffer* buffer);

// func (self Sound) SetBuffer(buffer SoundBuffer) { 
//     return C.sfSound_setBuffer(self.Cref, C.sfSound_setBuffer(buffer.Cref))
// }         

// \brief Get the audio buffer attached to a sound
// \param sound Sound object
// \return Sound buffer attached to the sound (can be NULL)
// const sfSoundBuffer* sfSound_getBuffer(const sfSound* sound);

// func (self SoundBuffer) GetBuffer(sound Sound) SoundBuffer { 
//     return SoundBuffer{C.sfSound_getBuffer(self.Cref)}	
// }

// \brief Set whether or not a sound should loop after reaching the end
// If set, the sound will restart from beginning after
// reaching the end and so on, until it is stopped or
// sfSound_setLoop(sound, sfFalse) is called.
// The default looping state for sounds is false.
// \param sound Sound object
// \param loop  sfTrue to play in loop, sfFalse to play once
// void sfSound_setLoop(sfSound* sound, sfBool loop);
func (self Sound) SetLoop(loop bool) {
	C.sfSound_setLoop(self.Cref, Bool(loop))
}

// \brief Tell whether or not a soud is in loop mode
// \param sound Sound object
// \return sfTrue if the sound is looping, sfFalse otherwise
// sfBool sfSound_getLoop(const sfSound* sound);
func (self Sound) GetLoop() bool {
	return C.sfSound_getLoop(self.Cref) == 1
}

// \brief Get the current status of a sound (stopped, paused, playing)
// \param sound Sound object
// \return Current status
// sfSoundStatus sfSound_getStatus(const sfSound* sound);
func (self Sound) Getstatus() SoundStatus { 
    switch C.sfSound_getStatus(self.Cref) {
	case 0: return Stopped
	case 1: return Paused
	case 2: return Playing
	}
	return UnknownSoundStatus
}

// \brief Set the pitch of a sound
// The pitch represents the perceived fundamental frequency
// of a sound; thus you can make a sound more acute or grave
// by changing its pitch. A side effect of changing the pitch
// is to modify the playing speed of the sound as well.
// The default value for the pitch is 1.
// \param sound Sound object
// \param pitch New pitch to apply to the sound
// void sfSound_setPitch(sfSound* sound, float pitch);
func (self Sound) SetPitch(pitch float32) {
	C.sfSound_setPitch(self.Cref, C.float(pitch))
}

// \brief Set the volume of a sound
// The volume is a value between 0 (mute) and 100 (full volume).
// The default value for the volume is 100.
// \param sound  Sound object
// \param volume Volume of the sound
// void sfSound_setVolume(sfSound* sound, float volume);
func (self Sound) SetVolume(volume float32) {
	C.sfSound_setVolume(self.Cref, C.float(volume))
}

// \brief Set the 3D position of a sound in the audio scene
// Only sounds with one channel (mono sounds) can be
// spatialized.
// The default position of a sound is (0, 0, 0).
// \param sound    Sound object
// \param position Position of the sound in the scene
// void sfSound_setPosition(sfSound* sound, sfVector3f position);
func (self Sound) SetPosition(v Vector3f) { 
    C.sfSound_setPosition(self.Cref, v.Cref)
}

/*
            
// \brief Make the sound's position relative to the listener or absolute
// Making a sound relative to the listener will ensure that it will always
// be played the same way regardless the position of the listener.
// This can be useful for non-spatialized sounds, sounds that are
// produced by the listener, or sounds attached to it.
// The default value is false (position is absolute).
// \param sound    Sound object
// \param relative sfTrue to set the position relative, sfFalse to set it absolute
// void sfSound_setRelativeToListener(sfSound* sound, sfBool relative);
func (self Sound) Setrelativetolistener(relative Bool) void { 
    return C.sfSound_setRelativeToListener(self.Cref, sfBool(relative));
}
*/

// \brief Set the minimum distance of a sound
// The "minimum distance" of a sound is the maximum
// distance at which it is heard at its maximum volume. Further
// than the minimum distance, it will start to fade out according
// to its attenuation factor. A value of 0 ("inside the head
// of the listener") is an invalid value and is forbidden.
// The default value of the minimum distance is 1.
// \param sound    Sound object
// \param distance New minimum distance of the sound
// void sfSound_setMinDistance(sfSound* sound, float distance);
func (self Sound) SetMinDistance(distance float32) {
	C.sfSound_setMinDistance(self.Cref, C.float(distance))
}

// \brief Set the attenuation factor of a sound
// The attenuation is a multiplicative factor which makes
// the sound more or less loud according to its distance
// from the listener. An attenuation of 0 will produce a
// non-attenuated sound, i.e. its volume will always be the same
// whether it is heard from near or from far. On the other hand,
// an attenuation value such as 100 will make the sound fade out
// very quickly as it gets further from the listener.
// The default value of the attenuation is 1.
// \param sound       Sound object
// \param attenuation New attenuation factor of the sound
// void sfSound_setAttenuation(sfSound* sound, float attenuation);

func (self Sound) Setattenuation(attenuation float32) {
	C.sfSound_setAttenuation(self.Cref, C.float(attenuation))
}

// \brief Change the current playing position of a sound
// The playing position can be changed when the sound is
// either paused or playing.
// \param sound      Sound object
// \param timeOffset New playing position
// void sfSound_setPlayingOffset(sfSound* sound, sfTime timeOffset);
func (self Sound) SetPlayingOffset(timeOffset Time) {
	C.sfSound_setPlayingOffset(self.Cref, timeOffset.Cref)
}

// \brief Get the pitch of a sound
// \param sound Sound object
// \return Pitch of the sound
// float sfSound_getPitch(const sfSound* sound);
func (self Sound) Getpitch() float32 {
	return float32(C.sfSound_getPitch(self.Cref))
}

// \brief Get the volume of a sound
// \param sound Sound object
// \return Volume of the sound, in the range [0, 100]
// float sfSound_getVolume(const sfSound* sound);
func (self Sound) GetVolume() float32 {
	return float32(C.sfSound_getVolume(self.Cref))
}

// \brief Get the 3D position of a sound in the audio scene
// \param sound Sound object
// \return Position of the sound in the world
// sfVector3f sfSound_getPosition(const sfSound* sound);
func (self Sound) Getposition() (v Vector3f) {
	p := C.sfSound_getPosition(self.Cref)
	return Vector3f{p}
}

// \brief Tell whether a sound's position is relative to the
//        listener or is absolute
// \param sound Sound object
// \return sfTrue if the position is relative, sfFalse if it's absolute
// sfBool sfSound_isRelativeToListener(const sfSound* sound);
func (self Sound) IsRelativeToListener() bool {
	return C.sfSound_isRelativeToListener(self.Cref) == 1
}

// \brief Get the minimum distance of a sound
// \param sound Sound object
// \return Minimum distance of the sound
// float sfSound_getMinDistance(const sfSound* sound);
func (self Sound) Getmindistance() float32 {
	return float32(C.sfSound_getMinDistance(self.Cref))
}

// \brief Get the attenuation factor of a sound
// \param sound Sound object
// \return Attenuation factor of the sound
// float sfSound_getAttenuation(const sfSound* sound);
func (self Sound) GetAttenuation() float32 {
	return float32(C.sfSound_getAttenuation(self.Cref))
}

// \brief Get the current playing position of a sound
// \param sound Sound object
// \return Current playing position
// sfTime sfSound_getPlayingOffset(const sfSound* sound);
func (self Sound) GetPlayingOffset() Time {
	return Time{C.sfSound_getPlayingOffset(self.Cref)}
}
