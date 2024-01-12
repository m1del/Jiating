import Quill from 'quill';
import BlotFormatter from "quill-blot-formatter";
import 'quill/dist/quill.snow.css';
import { useEffect } from "react";
import { useQuill } from "react-quilljs";

const ImageBlot = Quill.import('formats/image');

class CustomImageBlot extends ImageBlot {
  static create(value: any) {
    const node = super.create(value);
    // Add custom logic to manipulate the image node here
    node.setAttribute('data-is-display', 'false');
    return node;
  }
}

Quill.register(CustomImageBlot);


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

    ['image', 'video'],                                 //embeds

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
                const images = document.querySelectorAll('ql-editor img');
                images.forEach((img) => {
                    img.addEventListener('click', () => {
                        //logic to mark image as display image
                        // set data IsDisplay to true
                    })
                })

                // text is changing
                console.log('text change')
                console.log(delta);
                const currentContents = quill.getContents();
                console.log(currentContents.diff(oldContents));
            });
        }
    }, [quill, Quill]);


    return (
        <div className="p-4 bg-white shadow-lg rounded-lg max-w-4xl w-full">
            <div ref={quillRef} className="h-96" />
        </div>
    )
}

export default Editor;