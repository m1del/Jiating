import Quill from 'quill';
import BlotFormatter from "quill-blot-formatter";
import 'quill/dist/quill.snow.css';
import { useCallback, useEffect } from "react";
import { useQuill } from "react-quilljs";

const ImageBlot = Quill.import('formats/image');

/*

References: https://codesandbox.io/p/sandbox/react-quill-editor-with-image-resize-ry8vy?file=%2Fsrc%2FEditorWithUseQuill.js%3A10%2C17

*/

class CustomImageBlot extends ImageBlot {
  static create(value) {
    const outerNode = document.createElement('span'); // container for image and button
    outerNode.classList.add('custom-image-container');

    const node = super.create(value); // original image node
    outerNode.appendChild(node);

    const button = document.createElement('button');
    button.innerText = 'Set as Display Image';
    button.classList.add('toggle-display-btn', 'hidden', 'group-hover:block');
    outerNode.appendChild(button);

    // adjust button position based on image size
    // node.onload = () => { adjustButtonPosition(node, button) };

    return outerNode;
  }
}

const adjustButtonPosition = (img, button) => {
    const containerHeight = img.parentElement.offsetHeight;
    button.style.bottom = `${containerHeight - img.offsetHeight}px`;
}

Quill.register({
  'formats/image': CustomImageBlot,
  'modules/blotFormatter': BlotFormatter,
});

interface EditorProps {
    onContentChange: (content: string) => void;
}

const Editor: React.FC<EditorProps> = ({ onContentChange }) => {
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

    const { quill, quillRef } = useQuill({
        modules: { 
            toolbar: toolbarOptions,
            blotFormatter: {},
        }
    });

    const imageHandler = () => {
        const input = document.createElement('input');
        input.setAttribute('type', 'file');
        input.setAttribute('accept', 'image/*');
        input.click();

        input.onchange = async () => {
            const file = input.files[0];
            if (file) {
                // request presigned url from backend
                const presignedUrl = await getPresignedUrl(file);
                // upload image to s3 using presigned url
                const imgUrl = await uploadToS3(presignedUrl, file);

                // insert image in the editor and replace ref with s3 perma link
                const range = quill.getSelection();
                quill?.insertEmbed(range?.index, 'image', imgUrl);
            }
        }
    };

    // add custom image handler to toolbar
    useEffect(() => {
        if (quill) {
            quill.getModule('toolbar').addHandler('image', imageHandler);
        }
    }, [quill])

    // mock function to get presigned URL (TODO: needs implementation)
    async function getPresignedUrl(file) {
        console.log(file);
        console.log("get presigned url");
        await new Promise((resolve) => setTimeout(resolve, 1000));
        return 'https://www.google.com';
    }

    // mock function to upload file to S3 (TODO: needs implementation)
    async function uploadToS3(presignedUrl, event,file) {
        console.log("upload to s3")
        await new Promise((resolve) => setTimeout(resolve, 1000));
        return 'https://upload.wikimedia.org/wikipedia/commons/thumb/2/2f/Google_2015_logo.svg/640px-Google_2015_logo.svg.png';
    }

    const handleEditorClick = useCallback((e: React.MouseEvent<HTMLElement>) => {
        // check if click is outsade an image or its buttons
        if (e.target instanceof Element) {
            if (!e.target.classList.contains('toggle-display-btn') && e.target.tagName !== 'IMG') {
                // hide all buttons
                document.querySelectorAll('.toggle-display-btn').forEach(button => {
                    button.classList.add('hidden');
                })
            }
            // positioning button
            // if (e.target.tagName === 'IMG') {
            //     const img = e.target;
            //     const button = img.parentElement ? img.parentElement.querySelector('.toggle-display-btn') : null;
            //     adjustButtonPosition(img, button);
            // }
        }
    }, []);

    const handleToggleDisplayImage = useCallback((e) => {
        if (e.target.tagName === 'IMG') {
            const container = e.target.parentElement;
            const button = container.querySelector('.toggle-display-btn');
            button.classList.toggle('hidden');

            // update button based on current status
            const isDisplay = e.target.getAttribute('data-is-display');
            button.innerText = isDisplay === 'true' ? 'Unset as Display Image' : 'Set as Display Image';
        } else if (e.target.classList.contains('toggle-display-btn')) {
            const container = e.target.parentElement;
            const img = container.querySelector('img');

            // set or unset this image as display image
            const isDisplay = img.getAttribute('data-is-display') === 'true';

            if (!isDisplay) {
                // set image as the display iamge and set all other images as not display
                document.querySelectorAll('.ql-editor img').forEach(image => {
                    image.setAttribute('data-is-display', 'false');
                    const btn = img.parentElement ? img.parentElement.querySelector('.toggle-display-btn') : null;
                    btn.innerText = 'Set as Display Image';
                    btn?.classList.add("hidden");
                });
                img.setAttribute('data-is-display', 'true');
                e.target.innerText = 'Unset as Display Image';
            } else {
                // unset image as the display image
                img.setAttribute('data-is-display', 'false');
                e.target.innerText = 'Set as Display Image';
            }
        }
    }, []);
    
    useEffect(() => {
        if (quill) {
            const editor = quillRef.current;
            const contentContainer = editor.querySelector('.ql-editor');
            contentContainer.addEventListener('click', handleToggleDisplayImage);

            // hide all buttons on editor click
            contentContainer.addEventListener('click', handleEditorClick);

            return () => {
                contentContainer.removeEventListener('click', handleToggleDisplayImage);
                contentContainer.removeEventListener('click', handleEditorClick);
            }
        }
    }, [quill, quillRef, handleToggleDisplayImage, handleEditorClick]);

    // modified to allow parent component to access the editor content
    useEffect(() => {
        if (quill) {
            // combine the two event listeners into one lol idk
            quill.on('text-change', () => {
                const htmlContent = quill.root.innerHTML;
                onContentChange(htmlContent);
            })
        }
    }, [quill, onContentChange]);

    return (
        <div className="p-4 bg-white shadow-lg rounded-lg max-w-4xl w-full">
            <div ref={quillRef} className="h-96" />
        </div>
    )
}

export default Editor;