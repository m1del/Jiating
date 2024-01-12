import BlotFormatter from "quill-blot-formatter";
import 'quill/dist/quill.snow.css';
import { useEffect } from "react";
import { useQuill } from "react-quilljs";

const Editor = () => {
    const toolbarOptions = [
    ['bold', 'italic', 'underline', 'strike'],        // toggled buttons
    ['blockquote', 'code-block'],

    [{ 'header': 1 }, { 'header': 2 }],               // custom button values
    [{ 'list': 'ordered'}, { 'list': 'bullet' }],
    [{ 'script': 'sub'}, { 'script': 'super' }],      // superscript/subscript
    [{ 'indent': '-1'}, { 'indent': '+1' }],          // outdent/indent
    [{ 'direction': 'rtl' }],                         // text direction

    [{ 'size': ['small', false, 'large', 'huge'] }],  // custom dropdown
    [{ 'header': [1, 2, 3, 4, 5, 6, false] }],

    [{ 'color': [] }, { 'background': [] }],          // dropdown with defaults
    [{ 'font': [] }],
    [{ 'align': [] }],

    ['clean'],                                         // remove formatting
  ];

    const {quill, quillRef, Quill} = useQuill({
        modules: { 
            toolbar: toolbarOptions,
            blotFormatter: {} 
        }
    });

    if (Quill && !quill) {
        Quill.register('modules/blotFormatter', BlotFormatter);
    }

    useEffect(() => {
        if (quill) {
            quill.on('text-change', (delta, oldContents)  => {
                // text is changing
                console.log('text change')
                console.log(delta);
                const currentContents = quill.getContents();
                console.log(currentContents.diff(oldContents));
            });
        }
    }, [quill, Quill]);


    return (
        <div className="flex justify-center items-center h-screen bg-gray-100">
            <div className="editor-container p-4 bg-white shadow-lg rounded-lg max-w-4xl w-full">
                <div ref={quillRef} className="h-96" />
            </div>
        </div>
    )
}

export default Editor;