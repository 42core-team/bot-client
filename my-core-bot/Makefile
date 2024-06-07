SRCDIR = ./repo/src
OBJDIR = ./repo/build
SRCS = $(wildcard $(SRCDIR)/*.c)
OBJS = $(addprefix $(OBJDIR)/, $(notdir $(SRCS:.c=.o)))

CC = cc
CFLAGS = -Wall -Wextra -Werror
COREDIR = /core
LIBS = $(COREDIR)/con_lib.a
INC = $(COREDIR)
EXEC = bot

.PHONY: build clean fclean re $(EXEC)

build: fclean $(EXEC)

$(EXEC): $(OBJS)
	$(CC) $(CFLAGS) $^ -o $@ -L $(COREDIR) -l:con_lib.a -I $(INC)

$(OBJDIR)/%.o: $(SRCDIR)/%.c
	mkdir -p $(OBJDIR)
	$(CC) $(CFLAGS) -c $< -o $@ -I $(INC)

clean:
	rm -rf $(OBJDIR)

fclean: clean
	rm -rf $(EXEC)

re: fclean run
