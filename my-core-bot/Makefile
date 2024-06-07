SRCDIR = repo/src
OBJDIR = repo/obj
SRCS = $(wildcard $(SRCDIR)/*.c)
OBJS = $(addprefix $(OBJDIR)/, $(notdir $(SRCS:.c=.o)))

CC = cc
CFLAGS = -Wall -Wextra -Werror
CONDIR = ./connection
LIBS = $(CONDIR)/con_lib.a
INC = $(CONDIR)
EXEC = bot

.PHONY: build clean fclean re $(EXEC)

build: $(EXEC)
	mkdir -p $(OBJDIR)
	make -C $(CONDIR)

$(EXEC): $(OBJS)
	$(CC) $(CFLAGS) $^ -o $@ -L $(CONDIR) -l:con_lib.a -I $(INC)

$(OBJDIR)/%.o: $(SRCDIR)/%.c
	$(CC) $(CFLAGS) -c $< -o $@ -I $(INC)

clean: stop
	rm -rf $(OBJDIR)

fclean: clean
	rm -rf $(EXEC)

re: fclean run
