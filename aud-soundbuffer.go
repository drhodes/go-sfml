package sfml
// #include <SFML/Audio/Export.h>
// #include <SFML/Audio/Types.h>
// #include <SFML/Audio/SoundBuffer.h>
// #include <SFML/System/InputStream.h>
// #include <SFML/System/Time.h>
// #include <stddef.h>
import "C"

type SoundBuffer struct {
	Cref *C.sfSoundBuffer
}

// \brief Create a new sound buffer and load it from a file
// Here is a complete list of all the supported audio formats:
// ogg, wav, flac, aiff, au, raw, paf, svx, nist, voc, ircam,
// w64, mat4, mat5 pvf, htk, sds, avr, sd2, caf, wve, mpc2k, rf64.
// \param filename Path of the sound file to load
// \return A new sfSoundBuffer object (NULL if failed)
// sfSoundBuffer* sfSoundBuffer_createFromFile(const char* filename);
func SoundBufferFromFile(filename string) SoundBuffer { 
	fn := C.CString(filename)    
	return SoundBuffer{C.sfSoundBuffer_createFromFile(fn)}
}

// MEDITATE ON IF THIS IS FEASIBLE FROM GO
// \brief Create a new sound buffer and load it from a file in memory
// Here is a complete list of all the supported audio formats:
// ogg, wav, flac, aiff, au, raw, paf, svx, nist, voc, ircam,
// w64, mat4, mat5 pvf, htk, sds, avr, sd2, caf, wve, mpc2k, rf64.
// \param data        Pointer to the file data in memory
// \param sizeInBytes Size of the data to load, in bytes
// \return A new sfSoundBuffer object (NULL if failed)
// sfSoundBuffer* sfSoundBuffer_createFromMemory(const void* data, size_t sizeInBytes);
// func (self SoundBuffer) CreateFromMemory(sizeInBytes size_t) *SoundBuffer { 
//     return C.sfSoundBuffer_createFromMemory(self.Cref, sfSize_t(sizeInBytes));
// }

// MEDITATE ON IF THIS IS FEASIBLE FROM GO
// \brief Create a new sound buffer and load it from a custom stream
// Here is a complete list of all the supported audio formats:
// ogg, wav, flac, aiff, au, raw, paf, svx, nist, voc, ircam,
// w64, mat4, mat5 pvf, htk, sds, avr, sd2, caf, wve, mpc2k, rf64.
// \param stream Source stream to read from
// \return A new sfSoundBuffer object (NULL if failed)
// sfSoundBuffer* sfSoundBuffer_createFromStream(sfInputStream* stream);
// func (self SoundBuffer) Createfromstream() *SoundBuffer { 
//     return C.sfSoundBuffer_createFromStream(self.Cref);
// }
            
// \brief Create a new sound buffer and load it from an array of samples in memory
// The assumed format of the audio samples is 16 bits signed integer
// (sfInt16).
// \param samples      Pointer to the array of samples in memory
// \param sampleCount  Number of samples in the array
// \param channelCount Number of channels (1 = mono, 2 = stereo, ...)
// \param sampleRate   Sample rate (number of samples to play per second)
// \return A new sfSoundBuffer object (NULL if failed)
// sfSoundBuffer* sfSoundBuffer_createFromSamples(const sfInt16* samples, 
//       size_t sampleCount, unsigned int channelCount, unsigned int sampleRate);
// func (self SoundBuffer) CreateFromSamples(sampleCount size_t, 
// 	channelCount int32, sampleRate int32) SoundBuffer { 
//     return C.sfSoundBuffer_createFromSamples(self.Cref, sfSize_t(sampleCount), sf(int), sf(int));
// }
           
// \brief Create a new sound buffer by copying an existing one
// \param soundBuffer Sound buffer to copy
// \return A new sfSoundBuffer object which is a copy of \a soundBuffer
// sfSoundBuffer* sfSoundBuffer_copy(sfSoundBuffer* soundBuffer);
func (self SoundBuffer) Copy() SoundBuffer { 
    return SoundBuffer{C.sfSoundBuffer_copy(self.Cref)}
}
            
// \brief Destroy a sound buffer
// \param soundBuffer Sound buffer to destroy
// void sfSoundBuffer_destroy(sfSoundBuffer* soundBuffer);
func (self SoundBuffer) Destroy() { 
    C.sfSoundBuffer_destroy(self.Cref);
}

// \brief Save a sound buffer to an audio file
// Here is a complete list of all the supported audio formats:
// ogg, wav, flac, aiff, au, raw, paf, svx, nist, voc, ircam,
// w64, mat4, mat5 pvf, htk, sds, avr, sd2, caf, wve, mpc2k, rf64.
// \param soundBuffer Sound buffer object
// \param filename    Path of the sound file to write
// \return sfTrue if saving succeeded, sfFalse if it failed
// sfBool sfSoundBuffer_saveToFile(const sfSoundBuffer* soundBuffer, const char* filename);
func (self SoundBuffer) SaveToFile(filename string) bool { 
	fn := C.CString(filename)
    return C.sfSoundBuffer_saveToFile(self.Cref, fn) == 1
}

// TODO 
// \brief Get the array of audio samples stored in a sound buffer
// The format of the returned samples is 16 bits signed integer
// (sfInt16). The total number of samples in this array
// is given by the sfSoundBuffer_GetSampleCount function.
// \param soundBuffer Sound buffer object
// \return Read-only pointer to the array of sound samples
// const sfInt16* sfSoundBuffer_getSamples(const sfSoundBuffer* soundBuffer);
// func (self Soundbuffer) (SoundBuffer_getSamples)  { 
//     return C.sf*Int16(self.Cref, sf(*char));
// }

// TODO Verify how to handle size_t in this context.
// \brief Get the number of samples stored in a sound buffer
// The array of samples can be accessed with the
// sfSoundBuffer_getSamples function.
// \param soundBuffer Sound buffer object
// \return Number of samples
// size_t sfSoundBuffer_getSampleCount(const sfSoundBuffer* soundBuffer);
func (self SoundBuffer) GetSampleCount() C.size_t { 
    return C.sfSoundBuffer_getSampleCount(self.Cref);
}

// \brief Get the sample rate of a sound buffer
// The sample rate is the number of samples played per second.
// The higher, the better the quality (for example, 44100
// samples/s is CD quality).
// \param soundBuffer Sound buffer object
// \return Sample rate (number of samples per second)
// unsigned int sfSoundBuffer_getSampleRate(const sfSoundBuffer* soundBuffer);
func (self SoundBuffer) GetSampleRate() uint { 
	return uint(C.sfSoundBuffer_getSampleRate(self.Cref))
}
           
// \brief Get the number of channels used by a sound buffer
// If the sound is mono then the number of channels will
// be 1, 2 for stereo, etc.
// \param soundBuffer Sound buffer object
// \return Number of channels
// unsigned int sfSoundBuffer_getChannelCount(const sfSoundBuffer* soundBuffer);
func (self SoundBuffer) GetChannelCount() uint {
	return uint(C.sfSoundBuffer_getChannelCount(self.Cref))
}
           
// \brief Get the total duration of a sound buffer
// \param soundBuffer Sound buffer object
// \return Sound duration
// sfTime sfSoundBuffer_getDuration(const sfSoundBuffer* soundBuffer);
func (self SoundBuffer) Getduration() Time { 
    return Time{C.sfSoundBuffer_getDuration(self.Cref)}
}
            
