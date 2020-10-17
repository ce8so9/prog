%include 'asm4.s'
global main
extern exit

section .text
main:
  call read_hex
  mov edx,eax
  call read_hex
  add eax,edx
  add eax,eax
  inc eax

  call print_eax
  push 0
  call exit
