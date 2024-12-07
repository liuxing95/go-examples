# Makefile 示例

# 变量定义
CC = gcc
CFLAGS = -Wall -g
TARGET = my_program
SOURCES = main.c utils.c
OBJECTS = $(SOURCES:.c=.o)

# 默认目标
all: $(TARGET)

# 链接目标
$(TARGET): $(OBJECTS)
	$(CC) -o $@ $^

# 编译源文件
%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

# 清理生成的文件
clean:
	rm -f $(OBJECTS) $(TARGET)

.PHONY: all clean