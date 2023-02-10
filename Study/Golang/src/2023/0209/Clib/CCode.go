/*
 * @Author MierX
 * @Title CCode
 * @Date 2023-02-10 周一
 * @Time 14:03:28
 */

package Clib

/*
#include <stdio.h>
#include <windows.h>
#include <conio.h>

// 使用了WinAPI来移动控制台的光标
void gotoxy(int x, int y)
{
	COORD c;
	c.X=x,c.Y=y;
	SetConsoleCursorPosition(GetStdHandle(STD_OUTPUT_HANDLE),c);
}

// 从键盘获取一次按键，但不显示到控制台
int direct()
{
	return _getch();
}

// 去掉控制台光标隐藏
void hideCursor()
{
	CONSOLE_CURSOR_INFO cci;
	cci.bVisible = FALSE;
	cci.dwSize = sizeof(cci);
	SetConsoleCursorInfo(GetStdHandle(STD_OUTPUT_HANDLE), &cci);
}
*/
import "C"

// GotoPosition 设置控制台光标位置
func GotoPosition(x, y int) {
	C.gotoxy(C.int(x), C.int(y))
}

// Direction 无显示实现键盘输入的字符
func Direction() (key int) {
	key = int(C.direct())
	return
}

// HideCursor 设置控制台光标隐藏
func HideCursor() {
	C.hideCursor()
}
