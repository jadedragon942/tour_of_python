# all credit for this code goes to healeycodes
import sys
import resource
import pyseccomp as seccomp
import errno

MEMORY_LIMIT = 16 * 1024 * 1024  # 16 MB

CPU_TIME_LIMIT = 10  # 10sec
WRITE_LIMIT = 16 * 1024 * 1024  # 16 MB

def drop_perms():
    # Default action: return EPERM on disallowed syscalls
    filter = seccomp.SyscallFilter(seccomp.ERRNO(errno.EPERM))

    # Allow essential syscalls
    filter.add_rule(seccomp.ALLOW, "exit")
    filter.add_rule(seccomp.ALLOW, "exit_group")
    filter.add_rule(seccomp.ALLOW, "read")
    filter.add_rule(seccomp.ALLOW, "fstat")
    filter.add_rule(seccomp.ALLOW, "mmap")
    filter.add_rule(seccomp.ALLOW, "mprotect")
    filter.add_rule(seccomp.ALLOW, "munmap")
    filter.add_rule(seccomp.ALLOW, "brk")

    # Allow writing to stdout and stderr
    stdout_fd = sys.stdout.fileno()
    stderr_fd = sys.stderr.fileno()
    filter.add_rule(seccomp.ALLOW, "write", seccomp.Arg(0, seccomp.EQ, stdout_fd))
    filter.add_rule(seccomp.ALLOW, "write", seccomp.Arg(0, seccomp.EQ, stderr_fd))

    # Apply the filter
    filter.load()


def set_mem_limit():
    # virtual memory
    # resource.setrlimit(resource.RLIMIT_AS, (MEMORY_LIMIT, MEMORY_LIMIT))
    # cpu time
    resource.setrlimit(resource.RLIMIT_CPU, (CPU_TIME_LIMIT, CPU_TIME_LIMIT))
    # write limit i.e. don't allow an infinite stream to stdout/stderr
    resource.setrlimit(resource.RLIMIT_FSIZE, (WRITE_LIMIT, WRITE_LIMIT))


if __name__ == "__main__":
    codeFile = open(sys.argv[1], "r")
    code = codeFile.read()
    set_mem_limit()
    drop_perms()
    exec(code)
