#ifdef __cplusplus
extern "C" {
#endif

typedef void* GoSpoutSender;
GoSpoutSender GoCreateSender(const char* name, int width, int height);
bool GoSendTexture(GoSpoutSender s, unsigned int texture, int width, int height);

#ifdef __cplusplus
}
#endif
