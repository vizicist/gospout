#ifdef __cplusplus
extern "C" {
#endif

typedef void* GoSpoutSender;
GoSpoutSender GoCreateSender(const char* name, int width, int height);
bool GoSendTexture(GoSpoutSender s, unsigned int texture, int width, int height);
bool GoCreateReceiver(char* name, unsigned int* width, unsigned int* height, bool bUseActive);

#ifdef __cplusplus
}
#endif
