关键流程：
1. 音乐库功能, 使用者可以查看、添加、删除里面的音乐曲目
2. 播放音乐
3. 支持MP3, WAV, 但也能随时扩展以支持更多的音乐类型
4. 退出程序

程序运行后，进入一个循环，用于监听命令输入的状态，该程序接受一下命令。
1. 音乐库管理命令：lib, 包括list,add,remove
2. 播放管理功能：play命令，play后带歌曲名参数
3. 退出程序：q命令


Enter command -> 
lib add HugeStone MJ ~/MusicLib/hs.mp3 MP3
Enter command -> 
play HugeStone
Playing mp3 music ~/MusicLib/hs.mp3
..........
Finised playing ~/MusicLib/hs.mp3
Enter command -> 
lib list
1 : HugeStone ~/MusicLib/hs.mp3 MP3
Enter command -> 
lib view
Unrecognized lib command:  view
Enter command -> 
q