#ifdef __cplusplus
extern "C" {
#endif

typedef void* GoSpoutSender;
GoSpoutSender GoCreateSender(const char* name, int width, int height);
bool GoSendTexture(GoSpoutSender s, unsigned int texture, int width, int height);
bool GoCreateReceiver(char* name, unsigned int* width, unsigned int* height, bool bUseActive);
bool GoReceiveTexture(char* name, unsigned int* width, unsigned int *height, int textureID, int textureTarget, bool bInvert, int hostFBO);

#ifdef __cplusplus
}
#endif
