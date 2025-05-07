// Data structures
PageTable parent_page_table, child_page_table;
PhysicalPage *physical_pages[MAX_PAGES]; // array of physical page pointers
RefCount ref_count[MAX_PAGES];

// Function to fork a process using Copy-on-Write
Process* fork_process(Process *parent) {
    Process *child = create_process();
    
    for (int i = 0; i < MAX_PAGES; i++) {
        if (parent->page_table[i].valid) {
            // Point to same physical page
            child->page_table[i].frame = parent->page_table[i].frame;
            child->page_table[i].valid = true;
            child->page_table[i].writable = false;
            parent->page_table[i].writable = false;

            // Increment reference count
            ref_count[parent->page_table[i].frame]++;
        }
    }
    return child;
}

// On page fault (write to read-only shared page)
void handle_page_fault(Process *proc, int page_num) {
    int frame = proc->page_table[page_num].frame;
    
    if (ref_count[frame] > 1) {
        // Allocate new page
        int new_frame = alloc_physical_page();
        copy_page_contents(new_frame, frame);
        ref_count[frame]--;
        ref_count[new_frame] = 1;
        proc->page_table[page_num].frame = new_frame;
        proc->page_table[page_num].writable = true;
    } else {
        // Only this process owns the page now
        proc->page_table[page_num].writable = true;
    }
}

// Page deallocation (optional in this HW)
void free_physical_pages(Process *proc) {
    for (int i = 0; i < MAX_PAGES; i++) {
        if (proc->page_table[i].valid) {
            int frame = proc->page_table[i].frame;
            ref_count[frame]--;
            if (ref_count[frame] == 0) {
                dealloc_physical_page(frame);
            }
        }
    }
}
