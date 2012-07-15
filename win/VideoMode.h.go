////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////
/// \brief sfVideoMode defines a video mode (width, height, bpp, frequency)
///        and provides functions for getting modes supported
///        by the display device
///
////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////
/// \brief Get the current desktop video mode
///
/// \return Current desktop video mode
///
////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////
/// \brief Retrieve all the video modes supported in fullscreen mode
///
/// When creating a fullscreen window, the video mode is restricted
/// to be compatible with what the graphics driver and monitor
/// support. This function returns the complete list of all video
/// modes that can be used in fullscreen mode.
/// The returned array is sorted from best to worst, so that
/// the first element will always give the best mode (higher
/// width, height and bits-per-pixel).
///
/// \param count Pointer to a variable that will be filled with the number of modes in the array
///
/// \return Pointer to an array containing all the supported fullscreen modes
///
////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////
/// \brief Tell whether or not a video mode is valid
///
/// The validity of video modes is only relevant when using
/// fullscreen windows; otherwise any video mode can be used
/// with no restriction.
///
/// \param mode Video mode
///
/// \return sfTrue if the video mode is valid for fullscreen mode
///
////////////////////////////////////////////////////////////
