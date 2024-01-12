import BlotFormatter from "quill-blot-formatter";
import 'quill/dist/quill.snow.css';
import { useEffect } from "react";
import { useQuill } from "react-quilljs";

const Editor = () => {
    const {quill, quillRef, Quill} = useQuill({
        modules: { blotFormatter: {} }
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