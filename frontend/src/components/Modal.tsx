import React from 'react';

type ModalProps = {
  msg: string;
  onClose: () => void;
};

const Modal: React.FC<ModalProps> = ({ msg, onClose }) => {
  return (
    <div className="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full">
      {/* Modal Content */}
      <div className="relative top-20 mx-auto p-5 border w-96 shadow-lg rounded-md bg-white">
        <div className="mt-3 text-center">
          <h3 className="text-lg leading-6 font-medium text-gray-900">{msg}</h3>
          <div className="mt-2 px-7 py-3">
            <button
              onClick={onClose}
              className="bg-cyan-700 text-white font-bold py-2 px-4 rounded"
              type="button"
              style={{ transition: "all .15s ease" }}
            >
              Close
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Modal;
