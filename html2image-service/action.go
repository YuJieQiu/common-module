package html2image_service

func (r *ImageOptions) BuildImageRenderBytes() ([]byte, error) {

	out, err := GenerateImage(r)

	if err != nil {
		return nil, err
	}

	return out, nil
}
