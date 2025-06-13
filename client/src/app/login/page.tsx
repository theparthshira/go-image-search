"use client";
import React, { useEffect, useRef, useState } from "react";

type IImageList = {
  id: number;
  url: string;
};

export default function Page() {
  const [photos, setPhotos] = useState<IImageList[]>([]);
  const [loading, setLoading] = useState(false);
  const containerRef = useRef(null);

  // Mock API call
  const fetchPhotos = async (start = 0, count = 10): Promise<IImageList[]> => {
    const newPhotos: IImageList[] = Array.from(
      { length: count },
      (_, index) => ({
        id: start + index,
        url: `https://via.placeholder.com/150?text=Photo+${start + index + 1}`,
      })
    ).reverse();
    return new Promise((resolve) => setTimeout(() => resolve(newPhotos), 500));
  };

  // Load initial photos
  useEffect(() => {
    loadMorePhotos();
  }, []);

  const loadMorePhotos = async () => {
    if (loading) return;
    setLoading(true);

    const newPhotos = await fetchPhotos(photos.length, 10);

    console.log("newPhotos =====", newPhotos);

    // Maintain scroll position
    if (containerRef.current) {
      const element: HTMLElement = containerRef.current;
      const currentScrollHeight = element.scrollHeight;
      setPhotos((prev) => [...newPhotos, ...prev]);

      setTimeout(() => {
        element.scrollTop += element.scrollHeight - currentScrollHeight;
      }, 0);

      setLoading(false);
    }
  };

  console.log("photos =====", photos);

  const handleScroll = () => {
    if (containerRef.current) {
      const container: HTMLElement = containerRef.current;

      if (container.scrollTop === 0 && !loading) {
        loadMorePhotos();
      }
    }
  };

  return (
    <div
      ref={containerRef}
      onScroll={handleScroll}
      style={{
        height: "500px",
        overflow: "auto",
        display: "flex",
        flexDirection: "column",
        border: "1px solid #ccc",
      }}
    >
      {photos.map((photo) => (
        <div
          key={photo.id}
          style={{
            padding: "10px",
            borderBottom: "1px solid #ddd",
            textAlign: "center",
          }}
        >
          <img src={photo.url} alt={`Photo ${photo.id}`} />
        </div>
      ))}
      {loading && <div style={{ textAlign: "center" }}>Loading...</div>}
    </div>
  );
}
