SRCDIR = src
OBJDIR = build
SRCS = $(wildcard $(SRCDIR)/*.c)
OBJS = $(addprefix $(OBJDIR)/, $(notdir $(SRCS:.c=.o)))

CC = cc
CFLAGS = -Wall -Wextra -Werror
COREDIR = /core
LIBS = $(COREDIR)/con_lib.a
INC = $(COREDIR)
EXEC = bot

.PHONY: build clean fclean re $(EXEC)

build: $(EXEC)

$(EXEC): $(OBJS)
	$(CC) $(CFLAGS) $^ -o $@ -L $(COREDIR) -l:con_lib.a -I $(INC)

$(OBJDIR)/%.o: $(SRCDIR)/%.c
	$(CC) $(CFLAGS) -c $< -o $@ -I $(INC)

build: stop
	mkdir -p $(OBJDIR)

clean: stop
	rm -rf $(OBJDIR)

fclean: clean
	rm -rf $(EXEC)

re: fclean run
