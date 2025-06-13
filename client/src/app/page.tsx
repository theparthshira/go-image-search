"use client";
import Image from "next/image";
import { useEffect, useRef, useState } from "react";

export default function Home() {
  const [imageArr, setImageArr] = useState(
    Array(48)
      .fill(0)
      .map((_, i) => i + 1)
      .reverse()
  );

  console.log("imageArr =====", imageArr);

  const containerRef = useRef<null | HTMLDivElement>(null);
  const isScrolling = useRef(false); // To prevent redundant renders

  useEffect(() => {
    // Scroll to the bottom initially
    containerRef.current?.scrollTo(0, containerRef.current.scrollHeight);
  }, []);

  const handleScroll = (ele: HTMLElement) => {
    // if (isScrolling.current) return; // Avoid redundant updates
    // isScrolling.current = true;

    // requestAnimationFrame(() => {
    //   if (ele.scrollTop <= 400) {
    //     const currentHeight = ele.scrollHeight;

    //     setImageArr((state) => {
    //       const lastImage = state.length || 0;
    //       const newImages = Array(lastImage + 9)
    //         .fill(0)
    //         .map((_, i) => i + 1)
    //         .reverse();

    //       return newImages;
    //     });

    //     setTimeout(() => {
    //       const newHeight = ele.scrollHeight;

    //       console.log(
    //         "test =====",
    //         ele.scrollTop + (newHeight - currentHeight)
    //       );

    //       ele.scrollTop = ele.scrollTop + (newHeight - currentHeight) + 230;
    //       isScrolling.current = false;
    //     }, 0);
    //   } else {
    //     isScrolling.current = false;
    //   }
    // });

    if (ele.scrollTop < 150) {
      console.log("called =====");
      setImageArr((prev) => {
        const size = prev.length;
        const newArr = Array(45)
          .fill(0)
          .map((_, i) => size + i + 1)
          .reverse();

        console.log("newArr =====", newArr);
        console.log("prev =====", prev);

        return [...newArr, ...prev];
      });
      // useInterval(() => {

      if (containerRef.current) containerRef.current.scrollTop = 720;
    }
  };

  return (
    <div
      ref={containerRef}
      className="hideScrollbar max-w-[1296px] m-auto overflow-y-auto h-screen"
      onScroll={(e) => handleScroll(e.target as HTMLElement)}
    >
      <div className="flex justify-center">
        <div id="photos-all" className="flex flex-wrap">
          {imageArr?.map((image, idx) => (
            <div key={`test-${idx}`}>
              <Image
                key={idx}
                width={144}
                height={144}
                className="w-36 h-36 object-cover"
                alt={`Image ${image}`}
                src={`/${image}.jpg`}
              />
            </div>
          ))}
        </div>
      </div>
      <div
        style={{
          filter: "drop-shadow(5px 5px 42px #000000)",
          marginBottom: "250px",
        }}
      >
        <div className="text-4xl font-bold">Photos</div>
        <div className="text-base font-bold">2,375 Items</div>
      </div>
    </div>
  );
}
