export const error = $state({
    message: '',
    stack: '',
    reload_required: false,
    restart_required: false,

    show: false,
});

export function err(message: string, stack: string, reload_required: boolean = false, restart_required: boolean = false) {
    error.message = message;
    error.stack = stack;
    error.reload_required = reload_required;
    error.restart_required = restart_required;
    
    error.show = true;
}