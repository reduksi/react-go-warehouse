'use client';
import { useState, useRef, useEffect } from 'react';
import { mc } from '../../lib/mc';

export default function Select({
  value,
  name,
  label,
  onChange = () => {},
  className,
  required,
  options,
  readOnly,
  noDefault = true,
}) {
  const containerRef = useRef(null);
  const [selectedOption, setSelectedOption] = useState(value || '');
  const [isOpen, setIsOpen] = useState(false);

  const toggleOpen = () => setIsOpen(!isOpen);
  const handleOptionClick = (option) => {
    onChange({
      target: {
        value: option,
        name: name || '',
        checked: false,
        type: 'select',
      },
    });
    setSelectedOption(option);
    setIsOpen(false);
  };

  useEffect(() => {
    if (value !== undefined && value !== null) {
      setSelectedOption(String(value));
    } else if (!noDefault) {
      setSelectedOption('');
      onChange({
        target: {
          value: '',
          name: name || '',
          checked: false,
          type: 'select',
        },
      });
    }
    // eslint-disable-next-line
  }, [value]);

  useEffect(() => {
    const handleClickOutside = (event) => {
      if (
        containerRef.current &&
        !containerRef.current.contains(event.target)
      ) {
        setIsOpen(false);
      }
    };

    document.addEventListener('mousedown', handleClickOutside);
    return () => {
      document.removeEventListener('mousedown', handleClickOutside);
    };
  }, []);

  return (
    <main className={mc(`relative`, className)} ref={containerRef}>
      <div
        className={mc(
          `relative border rounded-md p-3 border-neutral-300`,
          `${selectedOption ? 'pt-6 pb-1' : 'pt-8 pb-4'}${
            readOnly
              ? ' bg-neutral-100 cursor-not-allowed'
              : ' focus:border-secondary-500'
          }`
        )}
        data-field-name={name}
        onClick={toggleOpen}
        tabIndex={readOnly ? -1 : 0}
      >
        <label
          className={mc(
            `absolute left-2 px-1 ${
              selectedOption
                ? 'text-xs text-neutral-500'
                : 'text-sm text-neutral-600'
            } transition-all duration-200 transform -translate-y-4 w-full overflow-hidden text-ellipsis whitespace-nowrap`
          )}
        >
          {label}
          {required && <span className="text-danger-500 ml-1 -mt-1">*</span>}
        </label>
        <div
          className={mc(
            'flex items-center justify-between',
            readOnly ? 'cursor-not-allowed' : 'cursor-pointer'
          )}
        >
          <span>{selectedOption}</span>{' '}
          <svg
            className={mc(`w-4 h-4 -mt-4`)}
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth={2}
              d="M19 9l-7 7-7-7"
            />
          </svg>
        </div>
      </div>
      {isOpen && !readOnly && (
        <div className="absolute mt-1 w-full bg-white rounded-md shadow-lg z-10">
          {options.map((option, index) => (
            <div
              key={index}
              className="px-4 py-2 cursor-pointer hover:bg-neutral-100"
              onClick={() => handleOptionClick(option)}
            >
              {option}
            </div>
          ))}
        </div>
      )}
    </main>
  );
}
