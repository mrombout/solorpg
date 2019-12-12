function! RollF()
    read !roll
    execute "normal! A\<cr>"
endfunction

command! Roll call RollF()