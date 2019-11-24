#include "CSpout_go.h"
#include "c:/Users/tjt/Documents/github/Spout2/SpoutSDK/Source/SpoutSender.h"
#include "c:/Users/tjt/Documents/github/Spout2/SpoutSDK/Source/SpoutReceiver.h"
#include <stdlib.h>

extern "C" {

// SpoutSender* S;
SpoutReceiver r;

GoSpoutSender GoCreateSender(const char* name, int width, int height) {
	SpoutSender* s = new SpoutSender;
	s->CreateSender(name,width,height);
	GoSpoutSender gss = (GoSpoutSender)s;
	return gss;
}

bool GoSendTexture(GoSpoutSender gss, unsigned int texture, int width, int height) {
	SpoutSender* s = (SpoutSender*)gss;
	return s->SendTexture(texture,GL_TEXTURE_2D,width,height);
}

bool GoCreateReceiver(char* name, unsigned int* width, unsigned int *height, bool bUseActive) {
	bool b;
	unsigned int w;
	unsigned int h;
	b = r.CreateReceiver(name,w,h,bUseActive);
	if ( b ) {
		*width = w;
		*height = h;
	}
	return b;
}

}
