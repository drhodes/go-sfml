package sfml
// #include <SFML/Audio/Export.h>
// #include <SFML/Audio/SoundBufferRecorder.h>
// #include <SFML/Audio/Types.h>
import "C"

type SoundBufferRecorder struct {
	Cref *C.sfSoundBufferRecorder
} 

// \brief Create a new sound buffer recorder
// \return A new sfSoundBufferRecorder object (NULL if failed)
// sfSoundBufferRecorder* sfSoundBufferRecorder_create(void);
func NewSoundBufferRecorder() SoundBufferRecorder { 
    return SoundBufferRecorder {
		C.sfSoundBufferRecorder_create(),
	}
}

// \brief Destroy a sound buffer recorder
// \param soundBufferRecorder Sound buffer recorder to destroy
// void sfSoundBufferRecorder_destroy(sfSoundBufferRecorder* soundBufferRecorder);
func (self SoundBufferRecorder) Destroy() { 
    C.sfSoundBufferRecorder_destroy(self.Cref)
}

           
// \brief Start the capture of a sound recorder recorder
// The \a sampleRate parameter defines the number of audio samples
// captured per second. The higher, the better the quality
// (for example, 44100 samples/sec is CD quality).
// This function uses its own thread so that it doesn't block
// the rest of the program while the capture runs.
// Please note that only one capture can happen at the same time.
// \param soundBufferRecorder Sound buffer recorder object
// \param sampleRate          Desired capture rate, in number of samples per second
// void sfSoundBufferRecorder_start(sfSoundBufferRecorder* soundBufferRecorder, unsigned int sampleRate);
func (self SoundBufferRecorder) Start(sampleRate uint) { 
    C.sfSoundBufferRecorder_start(self.Cref, C.uint(sampleRate))
}
            
// \brief Stop the capture of a sound recorder
// \param soundBufferRecorder Sound buffer recorder object
// void sfSoundBufferRecorder_stop(sfSoundBufferRecorder* soundBufferRecorder);
func (self SoundBufferRecorder) Stop() { 
    C.sfSoundBufferRecorder_stop(self.Cref);
}
           
// \brief Get the sample rate of a sound buffer recorder
// The sample rate defines the number of audio samples
// captured per second. The higher, the better the quality
// (for example, 44100 samples/sec is CD quality).
// \param soundBufferRecorder Sound buffer recorder object
// \return Sample rate, in samples per second
// unsigned int sfSoundBufferRecorder_getSampleRate(const sfSoundBufferRecorder* soundBufferRecorder);
func (self SoundBufferRecorder) GetSampleRate() uint {
	return uint(C.sfSoundBufferRecorder_getSampleRate(self.Cref))
}

// \brief Get the sound buffer containing the captured audio data
// The sound buffer is valid only after the capture has ended.
// This function provides a read-only access to the internal
// sound buffer, but it can be copied if you need to
// make any modification to it.
// \param soundBufferRecorder Sound buffer recorder object
// \return Read-only access to the sound buffer
// const sfSoundBuffer* sfSoundBufferRecorder_getBuffer(const sfSoundBufferRecorder* soundBufferRecorder);
func (self SoundBufferRecorder) GetBuffer() SoundBuffer { 
	return SoundBuffer{C.sfSoundBufferRecorder_getBuffer(self.Cref)}
}
