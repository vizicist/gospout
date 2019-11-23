#include "CSpout_go.h"
#include "/Users/tjt/Documents/github/Spout2/SpoutSDK/Source/SpoutSender.h"
#include <stdlib.h>

extern "C" {

// SpoutSender* S;

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

}
